type Default = {
  id: string;
  name: string;
};

type ComRuntime<Attr> = {
  data?: Record<string, any>;
  config: {
    attr: Attr;
  } & Default;
};

type GeneralComponentProps<Attr> = ComRuntime<Attr>;
