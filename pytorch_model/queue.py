import time
import stomp
import modelRunner

class ImageProcessorListener(stomp.ConnectionListener):
    def on_error(self, headers, message):
        print('received an error "%s"' % message)
    def on_message(self, headers, message):
        print(f'Received file name: {message}')
        modelRunner.process_image(message)
#connect to the activemq dockercontainer on the stomp port
conn = stomp.Connection(hosts_and_ports=[('activemq', 61613)])
#https://developers.redhat.com/blog/2018/06/14/stomp-with-activemq-artemis-python#setting_up_the_project
conn.set_listener('', ImageProcessorListener())
conn.start()
conn.connect('artemis', 'artemis', wait=True,headers = {'client-id': 'queuepy'} )

conn.subscribe(destination='queue.insect-finding-queue', id=1, ack='auto')
#Busy loop to keep the connection active and script running
try:
    while True:
        time.sleep(10)
except KeyboardInterrupt:
    print("Exiting")