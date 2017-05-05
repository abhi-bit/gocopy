### Feature set to support:

* Fast transfers
* Network error handling(connection/read timeouts)
* Checksum verification at source and destination for all transfers
* File compression. Prime motivation is to transfer db files, where
  blocks compress really well. Same cases observed where compression
  ratio was over 100x
* File permissions should remain intact
* Cross platform support

### Later enhancements:

* Transfer resume
* Bandwidth throttling
