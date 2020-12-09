import { atom, RecoilState } from 'recoil';


export const textState:RecoilState<string> = atom({
  key: 'textState',
  default: '',
});
