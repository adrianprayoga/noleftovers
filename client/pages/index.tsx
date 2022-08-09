import Head from "next/head";
import Layout, { siteTitle } from "../components/Layout";
import utilStyles from "../styles/utils.module.css";
import { getAllRecipes } from "../lib/recipes";
import RecipeGrid from "../components/RecipeGrid";
import Link from "next/link";

export default function Home({ allReciplesData }) {
  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <RecipeGrid recipeList={allReciplesData} />
      </section>
      <div className="flex flex-col justify-between p-4 leading-normal">
        <button className="mb-2 text-lg font-bold tracking-tight text-gray-900">
          <Link href="/recipe/create">Create Recipe</Link>
        </button>
      </div>
    </Layout>
  );
}

export async function getStaticProps() {
  const allReciplesData = await getAllRecipes();

  return {
    props: { allReciplesData },
  };
}
