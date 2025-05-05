import { Typography } from '@mui/material';
import React from 'react';

type FormLabelProps = {
  title: string | React.ReactNode;
  field: string;
};

const FormLabel: React.FC<FormLabelProps> = ({ title, field }) => (
  <Typography
    variant="subtitle1"
    fontWeight={600}
    component="label"
    htmlFor={field}
    mb="5px"
    mr="5px"
  >
    {title}
  </Typography>
);

export default FormLabel;
