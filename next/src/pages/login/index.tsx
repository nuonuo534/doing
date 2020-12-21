import React from 'react';
import { useRouter } from 'next/router';
import { useRecoilState } from 'recoil';
import { textState } from '@store/state';
import { Button, Form } from '@antd';
import BuildFormItem, {
  TypeColumnItem,
} from '@components/form/build_form_item';
import styles from './index.module.scss';

export default function Home() {
  const [text, setText] = useRecoilState(textState);
  const router = useRouter();
  const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value);
  };

  console.log(router);
  const goRouter = (event: React.MouseEvent<Element, MouseEvent>) => {
    event.preventDefault();
    router.push('/');
  };

  const columns: TypeColumnItem = [
    { label: '账号', name: 'username' },
    { label: '密码', name: 'password' },
  ];

  return (
    <div className={styles.page}>
      <Form>
        <BuildFormItem columns={columns} />
      </Form>
      <Button onClick={goRouter}>测试</Button>
    </div>
  );
}
