import { FC } from 'react';
import { Input } from '@antd';

interface TypeProps {
  typename: string;
  [P: string]: any;
}

const TypeComponent: FC<TypeProps> = (props) => {
  const typeName = props.typename?.toLowerCase();

  if (typeName === 'input') {
    return <Input {...props} />;
  } else {
    return <Input {...props} />;
  }
};

export default TypeComponent;
