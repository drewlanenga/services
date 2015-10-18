
github-services
==================

`github-services` is a collection of micro-services to run on [Google AppEngine](https://cloud.google.com/appengine/docs/go/).

# Installation/Setup

After getting an environment set up per the specs in the AppEngine [Go Tutorial](https://cloud.google.com/appengine/docs/go/gettingstarted/introduction), run:

```shell
go get -u -a github.com/drewlanenga/services

# for local development
cd $GOPATH/src/github.com/drewlanenga/services
goapp serve

# do it live!
cd $GOPATH/src/github.com/drewlanenga
goapp deploy -applicatiion <YOUR-PROJECT-ID> services/
```

# Usage

Development: http://localhost:8080/multibayes?text=Aaron%27s%20cat%20has%20fleas
Production: http://github-services.appspot.com/multibayes?text=Aaron%27s%20cat%20has%20fleas
