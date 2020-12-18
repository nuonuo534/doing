import { ReactNode } from 'react';
import { Input } from '@antd';

interface TypeReactNodeHash {
  [P: string]: ReactNode;
}

const keyComponent: TypeReactNodeHash = {
  input: Input,
};

export default function TypeComponent(name: string): ReactNode {
  const typeName = name.toLowerCase();
  const typeComponent = keyComponent[typeName];
  return typeComponent || null;
}
