import time
import stomp
import modelRunner

class ImageProcessorListener(stomp.ConnectionListener):
    def on_error(self, headers, message):
        print('received an error "%s"' % message)
    def on_connected(self, headers):
        print("Connected to the broker.")
    def on_message(self, *args, **kwargs):
        print(args[0].body)
        modelRunner.process_image(args[0].body)
print("Hey!")
#connect to the activemq dockercontainer on the stomp port
hosts = [('activemq', 61616)]
conn = stomp.Connection(host_and_ports=hosts, heartbeats=(20000, 0))
#https://developers.redhat.com/blog/2018/06/14/stomp-with-activemq-artemis-python#setting_up_the_project
conn.set_listener('', ImageProcessorListener())
conn.connect('artemis', 'artemis', wait=True,headers = {'client-id': 'queuepy'} )

conn.subscribe(destination='queue.insect-finding-queue', id=1, ack='auto')
#Busy loop to keep the connection active and script running
try:
    while True:
        time.sleep(10)
except KeyboardInterrupt:
    print("Exiting")