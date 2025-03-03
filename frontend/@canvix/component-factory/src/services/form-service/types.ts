export type FormUtils =
  | {
      registerField: (entity: any) => void;
      getFormValue: () => any;
      getFormKey: () => string;
      updateFormValue: (value: any, isEmitEvent?: boolean, path?: Array<string | number>) => void;
    }
  | Record<string, any>;

export type FormConfigType = {
  name: string;
  verifyEnable?: boolean;
  rules: Array<{
    type: 'system' | 'custom'; // 校验方式：内置函数/自定义函数
    message: string;
    system: string;
    custom: [
      // CFunction
      {
        id: string;
        content: string;
        name: string;
      },
    ];
  }>;
};
