import React from 'react';
import styles from './index.module.scss';

type Attr = {
  /** 文本方向 */
  // writingMode: string;
  /** 默认内容 */
  txt: string;
  /** 文字颜色，支持渐变色 */
  color: string;
  textStyle: {
    fontSize: number;
    fontFamily: string;
    lineHeight: number;
    fontWeight: string;
  };
  /** 最大显示行数 */
  lineClamp: number | null;
  textAlign: string;
  cursor: string;
  /** 字间距 */
  letterSpacing: number;
};

type Props = GeneralComponentProps<Attr, 'com'>;
const txtClass = styles['txt-component'];

export default class Txt extends React.Component<Props> {
  componentDidMount() {
    this.valueChanged();
  }

  componentDidUpdate(prevProps: Readonly<Props>) {
    this.handleValueChange(prevProps, this.props);
  }

  getData = (_props?: Props) => {
    const props = _props || this.props;
    const { source = [] } = props?.data || {};
    const data = Array.isArray(source) ? source : [];
    const value = props.config.attr?.txt;

    return data?.length > 0 ? { value, ...data[0] } : { value };
  };

  getValue = (): string => this.getData().value;

  getClassName = (attr: Attr) => {
    const { color, lineClamp } = attr;
    const isGradient = color.includes('gradient');
    return [
      txtClass,
      isGradient ? 'gradient-txt' : 'solid-txt',
      lineClamp ? 'line-clamp' : '',
    ].join(' ');
  };

  getStyle = (attr: Attr) => {
    const { textStyle, color, lineClamp, ...other } = attr;
    const { lineHeight, ...otherTextStyle } = textStyle;

    return {
      lineHeight: lineHeight ? `${lineHeight}px` : undefined,
      ...otherTextStyle,
      ...other,
      '--color': color,
      '--line-clamp': lineClamp,
    } as React.CSSProperties;
  };

  handleValueChange = (prevProps: Props, props: Props) => {
    const { source } = props.data;
    const { source: pSource } = prevProps.data;
    const { txt } = props.config.attr;
    const { txt: pTxt } = prevProps.config.attr;
    if (source !== pSource || txt !== pTxt) {
      this.valueChanged(props);
    }
  };

  valueChanged = (props?: Props) => {
    const params = this.getData(props);
    this.props.utils.$emit('valueChanged', params);
  };

  render(): React.ReactNode {
    const { attr } = this.props.config;
    const data = this.getValue();
    const style = this.getStyle(attr);
    const className = this.getClassName(attr);

    return (
      <div className={className} style={style}>
        {data}
      </div>
    );
  }
}
