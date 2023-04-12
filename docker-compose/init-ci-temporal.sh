#!/bin/bash

# use new version of tctl so that we can skip the prompt
tctl config set version next

for run in {1..60}; do
  sleep 1
  echo "now trying to register iWF system search attributes..."
  if tctl search-attribute  create -name IwfWorkflowType -type Keyword -y; then
    break
  fi
done

# see https://github.com/temporalio/temporal/issues/4160
sleep 0.1
tctl search-attribute  create -name IwfWorkflowType -type Keyword -y
sleep 0.1
tctl search-attribute  create -name IwfGlobalWorkflowVersion -type Int -y
sleep 0.1
tctl search-attribute  create -name IwfExecutingStateIds -type Keyword -y
sleep 0.1
tctl search-attribute  create -name CustomKeywordField -type Keyword -y
sleep 0.1
tctl search-attribute  create -name CustomIntField -type Int -y
sleep 0.1
tctl search-attribute  create -name CustomBoolField -type Bool -y
sleep 0.1
tctl search-attribute  create -name CustomDoubleField -type Double -y
sleep 0.1
tctl search-attribute  create -name CustomDatetimeField -type Datetime -y
sleep 0.1
tctl search-attribute  create -name CustomStringField -type text -y

tail -f /dev/null