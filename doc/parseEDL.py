from edl.edl import Parser
parser=Parser('23.98')
with open('Untitled Sequence.01.edl') as f:
    edl=parser.parse(f)
    for event in edl.events:
        #print "Event Number:"+str(event.num)
        #print "Source file:"+str(event.source_file)
        print "Clip Name:"+str(event.clip_name)
