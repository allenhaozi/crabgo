#! /bin/bash

echo "begin upgrade kratos \n"
helm upgrade iam-kratos ./charts/kratos --values values/kratos/values-dev.yaml -n openaios-iam

echo "begin upgrade hydra \n"
helm upgrade iam-hydra ./charts/hydra --values values/hydra/values-dev.yaml -n openaios-iam

