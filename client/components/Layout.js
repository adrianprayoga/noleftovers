import Head from 'next/head';
import Image from 'next/image';
import styles from '../styles/Layout.module.css';
import utilStyles from '../styles/utils.module.css';
import Link from 'next/link';
import Nav from './Nav';

const name = 'To change';
export const siteTitle = 'No Leftovers';

export default function Layout(props) {
  const { children, title } = props
  return (
    <>
      < Nav />
      <div className={styles.container}>
        <Head>
          <link rel="icon" href="/favicon.ico" />
          <meta
            name="noleftovers"
            content="Find recipes with ingredients that you have"
          />
          <meta
            property="og:image"
            content={`https://og-image.vercel.app/${encodeURI(
              siteTitle,
            )}.png?theme=light&md=0&fontSize=75px&images=https%3A%2F%2Fassets.vercel.com%2Fimage%2Fupload%2Ffront%2Fassets%2Fdesign%2Fnextjs-black-logo.svg`}
          />
          <meta name="og:title" content={siteTitle} />
          <meta name="twitter:card" content="summary_large_image" />
        </Head>
        <header className={styles.header}>
          <h1 className={utilStyles.heading2Xl}>{title}</h1>
        </header>
        <main>{children}</main>
        {/* {!home && (
          <div className={styles.backToHome}>
            <Link href="/">
              <a>← Back to home</a>
            </Link>
          </div>
        )} */}
      </div>
    </>
  );
}