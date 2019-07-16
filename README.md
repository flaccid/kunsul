# kunsul

Kubernetes global ingress UI dashboard.

## Installation

`go get github.com/flaccid/kunsul`

## Usage

`kunsul --help`

For example, `kunsul -D`

### Helm Chart

Validate the chart:

`helm lint charts/kunsul`

Install the chart:

`helm install --name kunsul charts/kunsul`

Upgrade the chart:

`helm upgrade kunsul charts/kunsul`

Upgrade chart with some custom values:

```
helm upgrade kunsul charts/kunsul \
  --recreate-pods \
  --set ingress.annotations."kubernetes\.io/ingress\.class"=nginx-internal \
  --set ingress.annotations."kubernetes\.io/ssl-redirect"=\"false\" \
  --set ingress.hosts[0]=kunsul.dev.my.cloud \
  --set image.pullPolicy=Always
```

Testing after deployment:

`helm test kunsul`

Completely remove the chart:

`helm delete --purge kunsul`

## Development

See all Make targets using `make help`.

For example, this runs a pod against the current kube context.

`make docker-build`
`make kube-test`

## License

- Author: Chris Fordham (<chris@fordham.id.au>)

```text
Copyright 2019, Chris Fordham

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
