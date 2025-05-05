type Default = {
  id: string;
  name: string;
};

type ComRuntime<Attr> = {
  data?: Record<string, any>;
  config: {
    attr: Attr;
  } & Default;
  utils: {
    $emit(eventName: string, payload: any): void;
  };
};

export type GeneralComponentProps<Attr> = ComRuntime<Attr>;
