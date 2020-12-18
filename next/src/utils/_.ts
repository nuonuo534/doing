import _ from 'lodash';


// export function omit<T extends object, K extends keyof T>(obj:T, ...keys: K[]):Omit<T, K> {
//   const newObj = {} as Omit<T, K>;
//   let key: keyof typeof obj;
//   for (key in obj) {
//       if (!keys.includes(key)) {
//         newObj[key] = obj[key];
//       }
//   }
//   return newObj;
// }