# sar
Search And Rescue - SAR

Given input data, split it into N parts, with M redundancy (initial implementation via RS code)
Encrypt each split portion with a different key
Upload to one or more S3 cloud services.

Retreive at least M parts from S3 cloud services, put them in order, decrypt each one with an appropriate key.
Recover original data from parts as retreived.

Enhancements to follow:
  * Compress each part, if smaller mark as compressed and use compressed data (potentially less data to up/download, reduced internal redundancy enhances encryption security.
  * Read only the minimum number of segments (n), no matter how many are available.  Only read additional if required to recover corrupted/bad data.
  * Conntrol files stored in a database?
  * File system more convenient for access.  
  * git back end?
  * connection to oss/cs key manager
  
