import RecipeCard from "./RecipeCard";

const RecipeGrid = (props) => {
  const { recipeList } = props;

  return (
    <>
      <div className={`p-10 grid gap-10 sm:md:grid-cols-1`}>
        {recipeList?.map(({ name, description, labels, id }, i) => (
          <RecipeCard key={i} id={id} name={name} description={description} />
        ))}
      </div>
    </>
  );
};

export default RecipeGrid;
