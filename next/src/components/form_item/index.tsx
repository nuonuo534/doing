import { FC, isValidElement, ReactElement } from 'react';
import getTypeComponent from './type_component';
// import { Form } from '@antd';

interface TypeColumnItem {
  render: () => ReactElement;
  [P: string]: any;
}

interface TypeFormItemProps {
  columns: [TypeColumnItem];
}

export const BuildFormItem: FC<TypeFormItemProps> = (props) => {
  const { columns } = props;
  return (
    <>
      {columns.map((item) => {
        let newComponent = null;
        const Render = item.render && item.render();
        if (Render && isValidElement(Render)) {
          newComponent = Render;
        } else {
          // const TypeComponent = getTypeComponent(item.type) as ReactElement;
          // newComponent = <TypeComponent {...(item.props || {})} />;
        }
      })}
    </>
  );
};
