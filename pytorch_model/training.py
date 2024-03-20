
import torchvision
from torch import nn, optim
from torch.utils.data import Dataset, DataLoader
from PIL import Image
import os
from torch.nn.functional import pad
import torch
from torchvision.models import ResNet152_Weights
from torchvision.models.detection import FasterRCNN
from torchvision.models.detection.backbone_utils import resnet_fpn_backbone
from torchvision.transforms import transforms

import sys
class AphidDamageDataset(Dataset):
    def __init__(self,  img_dir, transform=None,debugMode=False):

        self.img_dir = img_dir
        self.transform = transform
        self.debugMode=debugMode
        #read the ./training/parsed directory as images to train on
        self.unique_imgs = [f for f in os.listdir(img_dir) if os.path.isfile(os.path.join(img_dir, f))]

    def __len__(self):
        return len(self.unique_imgs)

    def __getitem__(self, idx):
        img_name = os.path.join(self.img_dir, self.unique_imgs[idx])
        image = Image.open(img_name).convert("RGB")


        width, height = image.size
        boxes = torch.tensor([[0, 0, width, height]], dtype=torch.float32)


        labels = torch.tensor([1], dtype=torch.int64)

        if self.transform:
            image = self.transform(image)
        target = {'boxes': boxes, 'labels': labels}

        return image, target


def collate_fn(batch):
    max_height = max(item[0].shape[1] for item in batch)
    max_width = max(item[0].shape[2] for item in batch)

    padded_images = []
    targets = []
    for image, target in batch:
        # Calculate padding
        height_pad = max_height - image.shape[1]
        width_pad = max_width - image.shape[2]

        # Pad the image
        padded_image = pad(image, (0, width_pad, 0, height_pad), "constant", 0)

        padded_images.append(padded_image)
        targets.append(target)

    return torch.stack(padded_images), targets


def main():
    if not torch.cuda.is_available():
        print("pytorch does not think cuda is available, this will be very slow on cpu taking over overnight or longer. You should set up cuda.")
        exit(-1)
    img_dir = './training/img/parsed'
    transform = transforms.Compose([
        transforms.ToTensor(),
        transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
    ])

    # Create an instance of the dataset
    aphid_dataset = AphidDamageDataset(img_dir=img_dir, transform=transform,debugMode=False)
    dataLoader = DataLoader(aphid_dataset, batch_size=10, shuffle=True, collate_fn=collate_fn)
    backbone = resnet_fpn_backbone(backbone_name='resnet152', weights=ResNet152_Weights.DEFAULT)
    model = FasterRCNN(backbone=backbone, num_classes=91)
    model.train()
    device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    model.to(device)
    optimizer = optim.SGD(model.parameters(), lr=0.005, momentum=0.9, weight_decay=0.0005)
    for epoch in range(0, 20):
        for images, targets in dataLoader:
            images = list(img.to(device) for img in images)

            targets = [{k: v.to(device) for k, v in t.items()} for t in targets]

            optimizer.zero_grad()

            loss_dict = model(images, targets)
            losses = sum(loss for loss in loss_dict.values())
            losses.backward()
            nn.utils.clip_grad_norm_(model.parameters(), max_norm=1)
            optimizer.step()
            print(f"Loss: {losses.item()}")
        print("looping")
    torch.save(model.state_dict(), "model.pth")

if __name__ == '__main__':
    main()
