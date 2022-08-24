import RecipeCard from "./RecipeCard";

const RecipeGrid = (props) => {
  const { recipeList } = props;

  return (
    <>
      <div className={`p-10 grid gap-10 sm:md:grid-cols-1`}>
        {recipeList?.map((recipe, i) => (
          <RecipeCard key={i} {...recipe} />
        ))}
      </div>
    </>
  );
};

export default RecipeGrid;
