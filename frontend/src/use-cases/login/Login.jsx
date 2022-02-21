import React, {useState} from "react";
import {
    FacebookLoginButton,
    GithubLoginButton,
    GoogleLoginButton,
    LoginButtonIcon,
    LoginButtonsContainer,
    LoginCard,
    LoginErrorText, MicrosoftLoginButton,
    StyledFacebookIcon
} from "./Login.styles";
import {getAuth} from "../../api/get.Auth.api";
import GitHubIcon from '@material-ui/icons/GitHub';
import {Typography} from "@material-ui/core";


const Login = () => {
    const [error, setError] = useState("");

    return (
        <div>
            <LoginCard>
                <Typography variant="h6">
                    Du behöver vara inloggad för att komma åt denna sida
                </Typography>
                {
                    error === "" ? (
                        <LoginButtonsContainer>
                            <GithubLoginButton onClick={() =>
                                login(setError)
                            }>
                                <LoginButtonIcon>
                                    <GitHubIcon/>
                                </LoginButtonIcon>
                                Logga in med github
                            </GithubLoginButton>
                        </LoginButtonsContainer>
                    ) : (
                        <LoginErrorText>
                            {error}
                        </LoginErrorText>
                    )
                }
            </LoginCard>
        </div>
    )
}

function login(setError) {
    getAuth()
        .then(response => {
            setError("Oväntat svar från servern")
        })
        .catch(error => {
            if (error.response && error.response.status === 401 && error.response.headers && error.response.headers.location) {
                window.location.assign(error.response.headers.location)
            } else {
                setError("Kunde inte logga in")
            }
        })
}

export default Login;