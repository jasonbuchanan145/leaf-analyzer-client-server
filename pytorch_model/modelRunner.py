import torchvision
from PIL import Image
import torchvision.transforms as transforms
import torch
import matplotlib.pyplot as plt
import matplotlib.patches as patches
from torchvision.models import ResNet152_Weights
from torchvision.models.detection import FasterRCNN
from torchvision.models.detection.backbone_utils import resnet_fpn_backbone

def _transform_image(image):
    transform = transforms.Compose([
        transforms.ToTensor(),
        transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
    ])
    return transform(image)

def _define_model():
    backbone = resnet_fpn_backbone(backbone_name='resnet152', weights=ResNet152_Weights.DEFAULT)
    model = FasterRCNN(backbone=backbone, num_classes=91)
    model = model.to('cuda')
    model.load_state_dict(torch.load('model.pth'))
    model.eval()
    return model
    
def _filter_and_draw(transformed_image,prediction,filename, image):
    pred_boxes = prediction[0]['boxes'].cpu().numpy()
    pred_labels = prediction[0]['labels'].cpu().numpy()
    pred_scores = prediction[0]['scores'].cpu().numpy()
    #the level of confidence in the score. anything below 50 is garbage
    threshold = 0.5
    filtered_boxes = pred_boxes[pred_scores >= threshold]
    filtered_labels = pred_labels[pred_scores >= threshold]
    filtered_scores = pred_scores[pred_scores >= threshold]

    #Load the image to draw the boxes over it
    pil_image = transforms.ToPILImage()(transformed_image.cuda())

    fig, ax = plt.subplots(1)
    ax.imshow(image)
    #Get the filtered boxes and their scores and overlay it with the image
    for box, score in zip(filtered_boxes,filtered_scores):
        x, y, x2, y2 = box
        rect = patches.Rectangle((x, y), x2 - x, y2 - y, linewidth=2, edgecolor='r', facecolor='none')
        ax.add_patch(rect)
        ax.text(x, y, f'{score:.2f}', fontsize=12, color='white', bbox=dict(facecolor='red', alpha=0.5))
    plt.savefig("./processed/"+filename.split('/')[-1].split(".")[0]+".png")

def process_image(filename):
    image = Image.open(filename).convert("RGB")
    transformed_image = _transform_image(image)
    model=_define_model()
    with torch.no_grad():
        prediction = model([transformed_image.to("cuda")])
    _filter_and_draw(transformed_image,prediction,filename,image)