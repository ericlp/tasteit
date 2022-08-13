import { useState } from "react";

import Image from "next/image";

import { Api, ApiResponse } from "../api/Api";
import { Button } from "../components/elements/Buttons/Buttons";
import { useTranslations } from "../hooks/useTranslations";
import CardLayout from "../layouts/CardLayout";

import styles from "./login.module.scss";

const Login = () => {
  const { t, translate } = useTranslations();
  const [error, setError] = useState<string | undefined>(undefined);

  const handleLoginResponse = (r: Promise<ApiResponse<unknown>>) => {
    r.then((val) => {
      if (val.error) {
        if (val.errorTranslationString) {
          setError(translate(val.errorTranslationString));
        } else {
          setError(t.errors.default);
        }
      }
    });
  };

  return (
    <CardLayout>
      <div className={`card ${styles.loginContainer}`}>
        <h1>{t.login.loginTitle}</h1>

        {error && (
          <div className="errorText marginTop marginBottom">{error}</div>
        )}

        <Button
          variant="primary"
          size="large"
          className={`${styles.signInButton} ${styles.microsoftButton}`}
          onClick={() => {
            handleLoginResponse(Api.user.login());
          }}
        >
          <div className={styles.signInButtonIcon}>
            <Image
              alt="Microsoft"
              src={"/itlogo.svg"}
              layout={"responsive"}
              width={"16px"}
              height={"16px"}
            />
          </div>
          {t.login.loginWithMicrosoft}
        </Button>

        <div className={"space"} />
      </div>
    </CardLayout>
  );
};

export default Login;
