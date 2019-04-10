# Golang - Domo API SDK

![Image of Domo](https://github.com/domoinc/domo-java-sdk/blob/master/domo.png)

## Warning

Domo does not officially support this golang API library

The original author doesn't either.  Consider it an orphan...

### About

* The Domo API SDK is the simplest way to automate your Domo instance
* The SDK streamlines the API programming experience, allowing you to significantly reduce your written code
* This package is a clone of the one previously published to [Jumbo Interactive Limited](https://github.com/JumboInteractiveLimited)

### Features

* DataSet and Personalized Data Policy (PDP) Management
  * Use DataSets for fairly static data sources that only require occasional updates via data replacement
  * Add Personalized Data Policies (PDPs) to DataSets (hide sensitive data from groups of users)
  * Docs: [Domo Developer Portal](https://developer.domo.com/docs/domo-apis/data)
* Stream Management
  * A Domo Stream is a specialized upload pipeline pointing to a single Domo DataSet
  * Use Streams for massive, constantly changing, or rapidly growing data sources
  * Streams support accelerated uploading via parallel data uploads
  * Docs: [Domo Developer Portal](https://developer.domo.com/docs/domo-apis/stream-apis)
* User Management
  * Create, update, and remove users
  * Major use case: LDAP/Active Directory synchronization
  * Docs: [Domo Developer Portal](https://developer.domo.com/docs/domo-apis/users)
* Group Management
  * Create, update, and remove groups of users
  * Docs: [Domo Developer Portal](https://developer.domo.com/docs/domo-apis/group-apis)

### Setup

go get:

```golang
go get github.com/davecb/domoStreamApi # was JumboInteractiveLimited/domostreamapi 
```

### Usage

* Create an API Client on the [Domo Developer Portal](https://developer.domo.com/)
* Use your API Client id/secret to instantiate a DomoClient()
* Multiple API Clients can be used by instantiating multiple Domo Clients
* Authentication with the Domo API is handled automatically by the SDK
* If you encounter a 'Not Allowed' error, this is a permissions issue. Please speak with your Domo Administrator.

### TODO
 - PageAPI is incomplete
 - Test coverage of user and group is poor
 - DataSet Export doesn't work reliably
 - In the StreamAPI there is a Modify() function that is not 100%
 - Need more examples
