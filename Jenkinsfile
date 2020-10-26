pipeline {  
  agent none
  environment {
        PATH = '/usr/local/bin' + $PATH 
        DB_ENGINE    = 'sqlite'
    }
  stages {
  	stage('Checkout code') {
        agent any
        steps {
            checkout scm            
        }
    }
  	stage('Docker Build') {
      agent any
      steps {      
        sh 'docker build -t system-api-img .'
      }
    }
    stage('Docker Run') {
      agent any
      steps {        
        sh 'docker run -p 9090:5000 --name system-api-con -e DbHost=host.docker.internal -e DbPort=5432 -e DbUser=logmaster -e  DbPassword=9psql%Ple1 -e DbName=events -d system-api-img'
      }
    }
  }
}