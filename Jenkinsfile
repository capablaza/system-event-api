pipeline {  
  agent none
  environment {
    APP_NAME =  "system-api"           
    IMAGE_NAME = "${env.APP_NAME}" + "-img"  
    CONTAINER_NAME = "${env.APP_NAME}" + "-con"
    OUT_CONTAINER_PORT="9090"
    IN_CONTAINER_PORT="5000"      
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
        echo "${env.PATH}"              
        sh "docker build -t ${IMAGE_NAME} -f Dockerfile ."
      }
    }
    stage('Docker Run') {
      agent any
      steps {                
        sh "docker run -p ${OUT_CONTAINER_PORT}:${IN_CONTAINER_PORT} --name ${CONTAINER_NAME} -e DbHost=host.docker.internal -e DbPort=5432 -e DbUser=logmaster -e  DbPassword=9psql%Ple1 -e DbName=events -d ${IMAGE_NAME} &"
      }
    }
  }
}