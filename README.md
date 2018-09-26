# transcode file structure
## Inbox
Put things here to be processed
## Outbox
Output of transcode goes here, possibly to be moved to plex automatically

# Redis Data Structures
## file queue
* general queue of files that have appeared in the inbox
## to-process queue
* queue of video files to be transcoded by workers
## processing store
* keep track of files being processed, and who they're being processed by.  mark as complete or remove when complete.

## Dispatcher
* monitor incoming file queue
* pop new file from file queue and push to to-process queue
* scale kubernetes deployment of workers up or down depending on queue

## API
* give visibility into queue length, etc
* interface with workers external to cluster to allow them to work the queue

## Worker
* pop file from to-process queue
* push file to being processed database (hash?)
* create configuration for `transcode-video`
* when transcode-video completes, mark as complete

## DTC CLI
* show information about queues and workers
* hit dispatcher API and request work
* create configuration for `transcode-video`
* when transcode-video completes, hit dispatcher API to mark as completed
