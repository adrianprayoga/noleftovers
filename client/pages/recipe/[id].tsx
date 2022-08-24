import Layout from "../../components/Layout";
import Head from "next/head";
import { GetStaticProps, GetStaticPaths } from "next";
import { getAllRecipeIds, getRecipeData } from "../../lib/recipes";
import utilStyles from "../../styles/utils.module.css";
import path from "path";

const basePath = "/images/recipe";

const Post = (props) => {
  const { recipeData } = props;

  return (
    <Layout home={false}>
      <h1 className="text-4xl font-extrabold leading-6 text-gray-900 my-4">
        {recipeData.name}
      </h1>
      {recipeData.imageLink && (
        <img
          className="h-96 rounded-lg md:h-auto md:w-48 mt-10"
          src={`${basePath}/${recipeData.imageLink}`}
          alt=""
        />
      )}
      <div className="text-gray-900 my-5">{recipeData.description}</div>
      <div className="bg-blue-50 p-5 rounded-md">
        <h1 className="text-2xl font-medium leading-6 text-gray-900 mb-4">
          Ingredients List
        </h1>
        <ul>
          {recipeData.ingredients?.map((ingredient) => {
            const ls = [
              ingredient.name,
              ingredient.amount,
              ingredient.measureValue?.String,
            ]
              .filter((e) => e)
              .join(" ");

            return (
              <li
                key={ingredient.position}
                className="text-gray-900 list-disc mx-5"
              >{`${ls}`}</li>
            );
          })}
        </ul>
      </div>
      <div className="p-5">
        <h1 className="text-2xl font-medium leading-6 text-gray-900 my-4">
          Steps
        </h1>
        {recipeData.steps?.map((step, i) => {
          return (
            <div key={i} className="my-5">
              <div className="font-bold">{`Step ${i + 1}`}</div>
              <p>{step.text}</p>
            </div>
          );
        })}
      </div>
    </Layout>
  );
};

export default Post;

export const getStaticPaths = async () => {
  // this needs to be known beforehand
  const paths = await getAllRecipeIds();
  return {
    paths,
    fallback: "blocking",
  };
};

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const recipeData = await getRecipeData(params.id);
  return {
    props: {
      recipeData,
    },
  };
};
