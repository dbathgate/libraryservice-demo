pipeline {
    agent any
    stages {
        stage('Build'){
            steps {
                script {
                    env.GOPATH = "${env.HOME}/go"
                }
                sh "make deps"
                sh "make build-linux"
            }
        }
        stage('Docker Image'){
            steps {
                sh "docker build -t library-service:${env.BUILD_ID} ."
            }
        }
        stage('Docker Push'){
            steps {
                sh "docker tag library-service:${env.BUILD_ID} localhost:5000/library-service:${env.BUILD_ID}"
                sh "docker push localhost:5000/library-service:${env.BUILD_ID}"
            }
        }

        stage ('Istio Initial Setup') {
            when {
                expression {
                    env.VERSION_COUNT = sh (
                        script: 'kubectl get deployments -l app=libraryservice -o name | wc -l',
                        returnStdout: true
                    ).trim()
                    return env.VERSION_COUNT.toInteger() == 0
                }
            }
            steps {
                sh "sed -i \"s/%%BUILD_NUMBER%%/${env.BUILD_ID}/g\" libraryservice-istio-route_template.yml"
                sh "kubectl apply -f libraryservice-istio-route_template.yml"
            }
        }

        stage('k8s Deploy') {
            steps {
                script {

                    env.PREVIOUS_VERSION_NUMBER = sh (
                        script: 'kubectl get deployments -l app=libraryservice -o json | jq ".items[].spec.template.metadata.labels.version" | sed "s/\\"//g"',
                        returnStdout: true
                    ).trim()

                    env.PREVIOUS_VERSION_NAME = sh (
                        script: "kubectl get deployments -l app=libraryservice,version=${env.PREVIOUS_VERSION_NUMBER} -o name",
                        returnStdout: true
                    ).trim()

                }

                sh "sed -i \"s/%%BUILD_NUMBER%%/${env.BUILD_ID}/g\" libraryservice-install_template.yml"
                sh "istioctl kube-inject -f libraryservice-install_template.yml > libraryservice-istio-install.yml"
                sh "kubectl apply -f libraryservice-service.yml"
                sh "kubectl apply -f libraryservice-istio-install.yml"
            }
        }

        stage('Canary') {
            when {
                expression {
                    env.VERSION_COUNT.toInteger() > 0
                }
            }
            steps {

                script {
                    def finishDeployment = false

                    while(!finishDeployment) {

                        userInput = input(
                            id: 'Proceed', message: 'Proceed Deployment?', parameters: [
                            [$class: 'TextParameterDefinition', defaultValue: "0", description: '', name: 'Increase Traffic?'],
                            [$class: 'BooleanParameterDefinition', defaultValue: false, description: '', name: 'Complete Deployment?']
                        ])

                        finishDeployment = userInput["Complete Deployment?"]
                        def greenPercentage = userInput["Increase Traffic?"].toInteger()
                        def bluePercentage = 100 - greenPercentage;

                        if (!finishDeployment) {
                            sh "cp libraryservice-canary_template.yml libraryservice-canary.yml"
                            sh "sed -i \"s/%%BLUE_VERSION%%/${env.PREVIOUS_VERSION_NUMBER}/g\" libraryservice-canary.yml"
                            sh "sed -i \"s/%%GREEN_VERSION%%/${env.BUILD_ID}/g\" libraryservice-canary.yml"
                            sh "sed -i \"s/%%BLUE_PERCENT%%/${bluePercentage}/g\" libraryservice-canary.yml"
                            sh "sed -i \"s/%%GREEN_PERCENT%%/${greenPercentage}/g\" libraryservice-canary.yml"
                            sh "kubectl apply -f libraryservice-canary.yml"
                        }
                    }
                }

                sh "sed -i \"s/%%BUILD_NUMBER%%/${env.BUILD_ID}/g\" libraryservice-istio-route_template.yml"
                sh "kubectl apply -f libraryservice-istio-route_template.yml"

                sh "kubectl delete ${env.PREVIOUS_VERSION_NAME}"
            }
        
        }
    }
    post { 
        always { 
            cleanWs()
        }
    }
}