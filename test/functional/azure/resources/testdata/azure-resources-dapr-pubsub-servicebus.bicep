param magpieimage string = 'radiusdev.azurecr.io/magpiego:latest' 

resource app 'radius.dev/Application@v1alpha3' = {
  name: 'azure-resources-dapr-pubsub-servicebus'

  resource publisher 'Container@v1alpha3' = {
    name: 'publisher'
    properties: {
      connections: {
        daprpubsub: {
          kind: 'dapr.io/PubSubTopic'
          source: pubsub.id
        }
      }
      container: {
        image: magpieimage
        env: {
          BINDING_DAPRPUBSUB_NAME: pubsub.name
          BINDING_DAPRPUBSUB_TOPIC: pubsub.properties.topic
        }
        readinessProbe:{
          kind:'httpGet'
          containerPort:3000
          path: '/healthz'
        }
      }
      traits: [
        {
          kind: 'dapr.io/Sidecar@v1alpha1'
          appId: 'publisher'
          appPort: 3000
        }
      ]
    }
  }
  
  resource pubsub 'dapr.io.PubSubTopic' = {
    name: 'pubsub'
    properties: {
      kind: 'pubsub.azure.servicebus'
      resource: namespace::topic.id
    }
  }
}

resource namespace 'Microsoft.ServiceBus/namespaces@2017-04-01' = {
  name: 'ns-${guid(resourceGroup().name)}'
  location: resourceGroup().location
  tags: {
    radiustest: 'azure-resources-dapr-pubsub-servicebus'
  }

  resource topic 'topics' = {
    name: 'TOPIC_A'
  }
}