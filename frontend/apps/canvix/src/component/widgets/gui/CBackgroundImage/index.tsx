import React from 'react';
import { isGradientColor } from '@canvas/utilities';
import { options } from './options';
import CBasic from '../CBasic';
import withGuiComponentWrapper from '../hoc/GuiComponentHoc';
import CSuite from '../CSuite';
import { getVisibleByConfig } from '../../PrivateProperty/utils';
import { ChangeParams, VisibleParams, ConfigType, CProps } from '@/types';

type Props = CProps<ConfigType>;
type State = {
  value: {
    type: 'image' | 'gradient';
    gradient: string;
    image: string | null;
  };
};

const initState = (value: string): State => {
  const { basic } = isGradientColor(value);
  return {
    value: {
      type: basic ? 'gradient' : 'image',
      gradient: basic ? value : '',
      image: basic ? '' : value
    }
  };
};

const getDefaultValue = (type: string) => (type === 'gradient'
  ? 'linear-gradient(40deg, rgb(27, 107, 235) 0%, rgb(25, 245, 157) 100%)'
  : null);

class CBackgroundImage extends CBasic<Props, State> {
  themeEnable = true;

  constructor(props: Props) {
    super(props);
    this.state = initState(props.value);
  }

  UNSAFE_componentWillReceiveProps(nextProps: Readonly<Props>): void {
    if (this.props.value !== nextProps.value) {
      this.setState(initState(nextProps.value));
    }
  }

  handleChange = (params: ChangeParams) => {
    const { keyChain, value } = params;
    const newChain = [...keyChain];
    const lastKey = newChain.pop() as unknown as keyof State['value'];
    const val = getDefaultValue(value);

    this.onChange({
      ...params,
      keyChain: newChain,
      value: lastKey === 'type' ? val : params.value,
      themeEnable: true
    });
  };

  getVisibleByConfig = (params: VisibleParams) => {
    const res = getVisibleByConfig({
      ...params,
      keyChain: [params.keyChain.at(-1)!],
      root: this.state.value
    });
    return res;
  };

  render(): React.ReactNode {
    const { keyChain, nameKeyChain, config } = this.props;
    const { themeEnable = true } = config;
    const { value } = this.state;

    return (
      <CSuite
        value={value}
        config={{ ...options, themeEnable }}
        utils={{ getVisibleByConfig: this.getVisibleByConfig }}
        level={0}
        keyChain={keyChain}
        nameKeyChain={nameKeyChain}
        onChange={this.handleChange}
      />
    );
  }
}

export default withGuiComponentWrapper(CBackgroundImage, { withoutLabel: true, themeEnable: true });
