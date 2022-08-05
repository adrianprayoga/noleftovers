import Head from 'next/head'
import Layout, { siteTitle } from '../components/Layout';
import utilStyles from '../styles/utils.module.css'
import { getAllRecipes } from '../lib/recipes';
import RecipeGrid from '../components/RecipeGrid';

export default function Home({ allReciplesData }) {
  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <RecipeGrid recipeList={allReciplesData} />
      </section>
    </Layout >
  );
}

export async function getStaticProps() {
  const allReciplesData = getAllRecipes()

  return {
    props: { allReciplesData }
  }
}