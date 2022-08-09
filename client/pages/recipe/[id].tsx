import Layout from "../../components/Layout";
import Head from "next/head";
import { GetStaticProps, GetStaticPaths } from "next";
import Date from "../../components/Date";
import { getAllRecipeIds, getRecipeData } from "../../lib/recipes";
import utilStyles from "../../styles/utils.module.css";
import path from "path";

const Post = (props) => {
  const { recipeData } = props;

  console.log(props)

  return (
    <Layout home={false}>
      {/* <Head>
        <title>{postData.title}</title>
      </Head>
      <h1 className={utilStyles.headingXl}>{postData.title}</h1>
      <div className={utilStyles.lightText}>
        <Date dateString={postData.date} />
      </div>
      <br />
      <div dangerouslySetInnerHTML={{ __html: postData.contentHtml }} /> */}
      <h1>{recipeData.name}</h1>
      <div>{recipeData.description}</div>
      
      
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
