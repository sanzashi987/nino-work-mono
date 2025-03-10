import React, { FC } from 'react';
import { ErrorOutline, Refresh } from '@mui/icons-material';
import ComStyleWrapper from './wrapper.module.scss';

type RuntimeErrorProps = {
  name?: string;
  retry?: () => void;
  overrideMessage?: string;
};

const { 'full-size': StyleWrapper } = ComStyleWrapper;

const errorClass = `${StyleWrapper} error frcc`;

const RuntimeError: FC<RuntimeErrorProps> = ({ name = '未知', retry, overrideMessage }) => (
  <div className={errorClass}>
    <ErrorOutline className="error" />
    {overrideMessage || `组件${name}发生错误`}
    {retry && (
      <span className="refresh">
        <Refresh onClick={retry} />
      </span>
    )}
  </div>
);

export default RuntimeError;
