/* eslint-disable react/no-unused-class-component-methods */
import React from 'react';
import type { CProps, PackageConfigType, ChangeParams } from '@/types';

class CBasic<Props extends CProps<PackageConfigType>, State = any> extends React.Component<Props, State> {
  themeEnable = false;

  onChange = (params: ChangeParams) => {
    const { onChange, config } = this.props;
    const { themeEnable = this.themeEnable, breakpointsEnable } = config;
    // 不在CBasic中统一判断的原因，当值不变，end由false转为true时会影响，如CColor,CSlider
    // if(this.props.value === params.value) return;
    onChange?.({
      themeEnable,
      breakpointsEnable,
      ...params
    });
  };

  render(): React.ReactNode {
    return null;
  }
}

export default CBasic;
