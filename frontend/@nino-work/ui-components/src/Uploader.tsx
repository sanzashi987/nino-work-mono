import { nanoid } from '@nino-work/shared';
import React, { createRef } from 'react';

export type Props = {
  accept?: string;
  multiple?: boolean;
  disabled?: boolean;
  children?: React.ReactNode;
  sizeLimit?: number;
  onChange?: (fileList: File[]) => void;
};

type State = {
  uid: string;
};

const UploaderContext = React.createContext<Uploader>(null);

class Uploader extends React.Component<Props, State> {
  inputRef = createRef<HTMLInputElement>();

  constructor(props: Props) {
    super(props);

    this.state = { uid: nanoid() };
  }

  // eslint-disable-next-line react/no-unused-class-component-methods
  open = () => {
    this.inputRef.current.click();
  };

  beforeUpload = (file: File): boolean => {
    const { accept, sizeLimit } = this.props;
    const { type, size } = file;
    const fileSize = size / 1024 / 1024;
    if (fileSize > sizeLimit) {
      return false;
    }

    const nameArr = file.name.split('.');
    const suffix = `.${nameArr[nameArr.length - 1]}`;
    if (accept) {
      let flag = false;
      const list = accept.split(',');
      for (let i = 0; i < list.length; i += 1) {
        const reg = new RegExp(list[i], 'gi');
        if (reg.test(suffix) || reg.test(type)) {
          flag = true;
          break;
        }
      }
      // MIME type && 后缀名均不包含，返回false
      if (!flag) {
        return false;
      }
    }
    // }
    return true;
  };

  onChange = (files: File[]) => {
    const { onChange } = this.props;
    let res: File[] = files.concat();
    res = res.filter(file => this.beforeUpload(file));
    onChange?.(res);
    this.reset();
  };

  handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { files } = e.target;
    if (!files) return;

    this.onChange(Array.from(files));
  };

  reset() {
    this.setState({ uid: nanoid() });
  }

  render() {
    const { disabled, accept, multiple, children } = this.props;
    const { uid } = this.state;

    return (
      <UploaderContext.Provider value={this}>
        <div onClickCapture={this.open}>
          <input
            key={uid}
            ref={this.inputRef}
            type="file"
            hidden
            disabled={disabled}
            accept={accept}
            multiple={multiple}
            onChange={this.handleChange}
          />
          {children}
        </div>
      </UploaderContext.Provider>
    );
  }
}
export class Droppable extends React.Component {
  static contextType = UploaderContext;

  declare context: Uploader;

  ref = createRef<HTMLDivElement>();

  state: { dragState: string } = { dragState: '' };

  handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.stopPropagation();
    e.preventDefault();
    this.setDragState(e);
  };

  setDragState = (e: React.DragEvent<HTMLDivElement>) => {
    const { type } = e;
    const { dragState } = this.state;
    if (type !== dragState) {
      this.setState({ dragState: type });
    }
  };

  handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.stopPropagation();
    e.preventDefault();
    this.setDragState(e);
    const { props, onChange } = this.context;
    let res: File[] = Array.from(e.dataTransfer.files);
    if (!props.multiple) {
      res = res.slice(0, 1);
    }
    onChange(res);
  };

  render(): React.ReactNode {
    const { open } = this.context;
    return (
      <div
        ref={this.ref}
        role="button"
        tabIndex={-1}
        onDrop={this.handleDrop}
        onClick={open}
        onDragOver={this.handleDragOver}
        onDragLeave={this.setDragState}
      />
    );
  }
}

export default Uploader;
