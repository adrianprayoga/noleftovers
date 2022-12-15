import RecipeCard from "./RecipeCard";

const RecipeGrid = (props) => {
  const { recipeList, favorites, handleAddFavorite, handleRemoveFavorite, userAuthenticated } = props;

  return (
    <>
      <div className={`p-10 grid gap-10 sm:md:grid-cols-1`}>
        {recipeList?.map((recipe, i) => (
          <RecipeCard
            key={i}
            isFavorite={favorites.indexOf(recipe.id) !== -1}
            isDisabled={!userAuthenticated}
            handleAddFavorite={handleAddFavorite}
            handleRemoveFavorite={handleRemoveFavorite}
            {...recipe}
          />
        ))}
      </div>
    </>
  );
};

export default RecipeGrid;
