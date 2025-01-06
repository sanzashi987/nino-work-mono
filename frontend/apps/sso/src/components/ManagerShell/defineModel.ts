type Model<T = string> = {
  label: string
  field:string
};

const defineModel = <T>(m: Model<T>): Model<T> => m;
