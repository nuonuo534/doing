import React from 'react';
import { AppProps } from 'next/app';
import { RecoilRoot } from 'recoil';
import { ConfigProvider, AntdConfigProps } from '@antd';
import '@styles/globals.scss';
import 'antd/dist/antd.css';

const MyApp: React.FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <RecoilRoot>
      <ConfigProvider {...AntdConfigProps}>
        <Component {...pageProps} />
      </ConfigProvider>
    </RecoilRoot>
  );
};

export default MyApp;
