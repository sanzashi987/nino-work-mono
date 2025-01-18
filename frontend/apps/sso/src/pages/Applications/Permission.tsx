import React from 'react';
import { useParams } from 'react-router-dom';

type PermissionManagementProps = {};

const PermissionManagement: React.FC<PermissionManagementProps> = (props) => {
  const { appId } = useParams();

  return appId;
};

export default PermissionManagement;
