import { FormEvent, useState } from "react";

import { Api } from "../../api/Api";
import { EDIT_RECIPE_BASE_ENDPOINT } from "../../api/Endpoints";
import { Button } from "../../components/elements/Buttons/Buttons";
import Dropdown from "../../components/elements/Dropdown/Dropdown";
import TextField from "../../components/elements/TextField/TextField";
import { useMe } from "../../hooks/useMe";
import { useTranslations } from "../../hooks/useTranslations";
import CardLayout from "../../layouts/CardLayout";

import styles from "./create.module.scss";

const CreateRecipe = () => {
  const { t, translate } = useTranslations();
  const [error, setError] = useState<string | undefined>(undefined);
  const [name, setName] = useState("");

  const { me } = useMe();

  const ownerOptions = (me?.owners ?? []).map((owner) => {
    return {
      display: owner.name,
      value: owner.id,
    };
  });

  const [ownerId, setOwnerId] = useState("");

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const owner = ownerId ? ownerId : ownerOptions[0].value;

    Api.recipes.create(name, owner).then((data) => {
      if (data.error && data.errorTranslationString) {
        setError(translate(data.errorTranslationString));
      } else {
        const uniqueName = data.data?.uniqueName;
        window.location.assign(`${EDIT_RECIPE_BASE_ENDPOINT}/${uniqueName}`);
      }
    });
  };

  return (
    <CardLayout>
      <form
        className={`card ${styles.createCardColumn}`}
        onSubmit={(e) => handleSubmit(e)}
      >
        <h3 className="marginBottom">{t.recipe.createRecipeTitle}</h3>
        <div className={`marginBottom ${styles.createElement}`}>
          <div>
            <label htmlFor="recipe_name" className="marginRight">
              {t.recipe.recipeName}
            </label>
            <TextField
              variant="outlined"
              name="recipe_name"
              id="recipe_name"
              focus={true}
              placeholder={t.recipe.recipeName}
              value={name}
              onChange={(e) => {
                setName(e.target.value);
              }}
              required
              className={styles.createTextField}
            />
          </div>
          {ownerOptions.length > 1 && (
            <>
              <label htmlFor="recipe_owner" className="marginRight">
                {t.recipe.recipeOwner}
              </label>
              <Dropdown
                name="recipe_owner"
                onUpdate={(ownerId) => setOwnerId(ownerId)}
                options={ownerOptions}
                defaultValue={ownerOptions[0].value}
                visibleSize={"auto"}
                variant={"outlined"}
              />
            </>
          )}
        </div>
        {error && <p className="errorText">{error}</p>}
        <Button
          variant="primary"
          size="normal"
          className={styles.createRecipeButton}
          type="submit"
        >
          {t.recipe.createRecipe}
        </Button>
      </form>
    </CardLayout>
  );
};

export default CreateRecipe;
