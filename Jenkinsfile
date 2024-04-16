pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout the repository
                git 'https://github.com/gocheto8/si_jenkins_test.git'
            }
        }
        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build -o si_jenkins/api'
            }
        }
        stage('Test') {
            steps {
                // Run tests
                sh 'go test ./...'
            }
        }
    }

    post {
        success {
            echo 'Build and tests succeeded!'
            // You can add additional steps here, like deploying the application
        }
        failure {
            echo 'Build or tests failed :('
            // You can add additional steps here, like sending notifications
        }
    }
}
