import { FC } from 'react';
import { Input } from '@antd';

interface TypeProps {
  typeName: string;
  [P: string]: any;
}

const TypeComponent: FC<TypeProps> = (props) => {
  const typeName = props.typeName.toLowerCase();
  if (typeName === 'input') {
    return <Input {...props} />;
  } else {
    return <Input {...props} />;
  }
};

export default TypeComponent;
