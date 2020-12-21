import { FC, isValidElement, ReactNode } from 'react';
import { Form } from '@antd';
import TypeComponent from './type_component';

export interface TypeColumnItem {
  label: string;
  name: string | number | (string | number)[];
  render?: () => ReactNode;
  type?: string;
  [P: string]: any;
}

interface TypeFormItemProps {
  columns: [TypeColumnItem];
}

const BuildFormItem: FC<TypeFormItemProps> = (props) => {
  const { columns } = props;
  return (
    <>
      {columns.map((item) => {
        let newComponent!: JSX.Element;
        const Render = item.render && item.render();
        if (Render && isValidElement(Render)) {
          newComponent = Render;
        } else {
          const newProps = { ...(item.props || {}), typeName: item.type };
          newComponent = <TypeComponent {...newProps} />;
        }
        const key = Array.isArray(item.name) ? item.name.join('-') : item.name;
        return (
          <Form.Item key={key} {...item}>
            {newComponent}
          </Form.Item>
        );
      })}
    </>
  );
};

export default BuildFormItem;
