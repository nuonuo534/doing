import Head from 'next/head';
import { useRecoilValue } from 'recoil';
import { textState } from '@store/state';
import styles from '@styles/Home.module.css';

export default function Home() {
  const text = useRecoilValue(textState);
  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      {text}
    </div>
  );
}
