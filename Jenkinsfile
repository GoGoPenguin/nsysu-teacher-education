pipeline {

    agent any

    environment {
        REPO = "https://github.com/SuperEdge/nsysu-teacher-education.git"
        NODEJS = "NodeJS 10.16.3"
        GOLANG = "C:\\Go\\bin\\go.exe"
        SQL_MIGRATE = "C:\\Users\\Administrator\\go\\bin\\sql-migrate.exe"
        NSSM = "C:\\nssm-2.24\\win64\\nssm.exe"
        API_SERVICE = "ctep-api"
    }

    stages{
        stage('Checkout') {
            steps {
                echo 'Checkout'

                script {
                    if (env.DEPLOY_BRANCH == "commit") {
                        DEPLOY_BRANCH = env.DEPLOY_VALUE
                    } else if (env.DEPLOY_BRANCH == "tag") {
                        DEPLOY_BRANCH = "refs/tags/" + env.DEPLOY_VALUE
                    } else if (env.DEPLOY_BRANCH == "branch") {
                        DEPLOY_BRANCH = "*/" + env.DEPLOY_VALUE
                    }
                }

                checkout([$class: 'GitSCM', branches: [[name: "${DEPLOY_BRANCH}"]],
                    userRemoteConfigs: [[url: "${REPO}"]],
                    extensions: [[$class: 'CloneOption', shallow: true]]])
            }
        }

        stage('Build') {
            steps {
                echo 'Build'

                script {
                    bat "cd api/db & ${SQL_MIGRATE} up -env=production"

                    bat "cd api/db/seeder & ${GOLANG} run ."

                    bat "cd api & ${GOLANG} build -o main.exe main.go"

                    nodejs(nodeJSInstallationName: "${NODEJS}") {
                        bat "cd front & yarn install"
                        bat "cd back & yarn install"
                    }
                }
            }
        }

        stage('Test') {
            steps {
                echo "Test";
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploy";

                script {
                    bat "${NSSM} restart ${API_SERVICE}"
                }
            }
        }
    }
}
