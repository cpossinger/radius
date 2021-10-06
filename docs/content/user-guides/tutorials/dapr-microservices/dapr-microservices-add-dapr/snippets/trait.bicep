//SAMPLE
resource app 'radius.dev/Application@v1alpha3' = {
  name: 'dapr-tutorial'

  resource backend 'ContainerComponent' = {
    name: 'backend'
    properties: {
      //RUN
      container: {
        image: 'radius.azurecr.io/daprtutorial-backend'
      }
      //RUN
      traits: [
        {
          kind: 'dapr.io/Sidecar@v1alpha1'
          appPort: 3000
        }
      ]
    }
  }
}
//SAMPLE
