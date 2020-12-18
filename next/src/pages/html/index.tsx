import React from 'react';
import { useRouter } from 'next/router';
import { useRecoilState } from 'recoil';
import { textState } from '@store/state';
import { Button } from '@antd';
import styles from '@styles/Home.module.scss';

export default function Home() {
  const [text, setText] = useRecoilState(textState);
  const router = useRouter();
  const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value);
  };

  const goRouter = (event: React.MouseEvent<Element, MouseEvent>) => {
    event.preventDefault();
    router.push('/');
  };

  return (
    <div className={styles.container}>
      <input type="text" value={text} onChange={onChange} />
      {text}
      <Button onClick={goRouter}>测试</Button>
    </div>
  );
}
