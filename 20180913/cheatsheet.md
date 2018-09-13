Cheat sheet
---

### C.1

#### env setup

~~~
# source <(oc completion bash)
~~~

~~~
# oc new-project $TODO_PROJECT_NAME
~~~

#### basic object

~~~
# oc explain

# oc explain pod
~~~

#### LAB: Deploy some objects

~~~
# oc create -f https://raw.githubusercontent.com/nak3/demo3/master/20180913/pod.yaml
~~~

~~~
# oc create -f https://raw.githubusercontent.com/nak3/demo3/master/20180913/dc.yaml
~~~

~~~
# oc create -f https://raw.githubusercontent.com/nak3/demo3/master/20180913/bc.yaml
~~~

#### LAB: template and oc new-app

~~~
# oc get template -n openshift
# oc new-app --template=eap70-basic-s2i
~~~

~~~
# oc new-app https://github.com/openshift/ruby-ex.git
~~~

### Ch.2

### Getting objects

~~~
# oc get all
~~~

~~~
# oc describe dc
~~~

~~~
# oc get dc -o yaml
~~~

### Getting log

~~~
# oc logs $TOD_POD_NAME 
~~~

~~~
# oc logs $TOD_POD_NAME -p
~~~

~~~
# oc get event
~~~



~~~
# oc rsh $TOD_POD_NAME 

- or -

# oc exec -it $TOD_POD_NAME bash
~~~

~~~
# oc debug $TOD_POD_NAME 
~~~

~~~
# oc rsync $TODO_POD_NAME:/tmp/ .
~~~

### Ch.3

~~~
# oc set env dc $TODO_DC_NAME AUTO_DEPLOY_EXPLODED=true
~~~

~~~
# oc set env dc $TODO_DC_NAME --list
~~~

~~~
# oc create configmap --from-literal=testkey=foo  envcmp
~~~

~~~
# echo "configmap test" > test.txt

# oc create configmap --from-file=test.txt  filecmp
~~~


~~~
# oc set env --from=configmap/envcmp  dc/$TODO_DC_NAME
~~~

~~~
# oc set volumes dc/$TODO_DC_NAME --add --name=testcmp \
     -t configmap  --configmap-name=filecmp   --mount-path=/mnt/
~~~
