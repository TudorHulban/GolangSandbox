# ScreenshotURL
Take screenshots of URLs. No time stamp on screenshot (see before and after).

### Requirements
#### 1. SaaS
#### 2. Inputs:
##### a. URL (JSON)
##### b. URLs (JSON array)
#### 3. Output:
##### a. URL of image (JSON)
##### b. URLs of images (JSON array) 

### Solution
##### 1. Workers running in Alpine (small footprint) containers, receiving URLs from a Dispatcher
##### 2. Each Worker runs one Firefox process
##### 3. Workers save images on shared NFS
##### 4. NFS shares provided by external storage are syncronized with Rsync
##### 5. Dispatcher is the end point for sending the requests
##### 6. No service discovery, Dispatcher knows the IPs of Workers
##### 7. Communication between Dispatcher and Worker is done by TCP
##### 8. Dispatcher serves a limited number of requests. Redirection to another Dispatcher when the limit is reached

### Code Contains
##### a. Worker
##### b. Client to test Worker

### Assumptions
##### 1. No error validation for URL, ex. 404
##### 2. gRPC no additional value in bringing it in
