import React from 'react';
import { useParams } from 'react-router-dom';
import { listPermissions } from '@/api';
import { usePromise } from '@/utils';
import loading from '@/components/Loading';

type PermissionManagementProps = {};

const PermissionManagement: React.FC<PermissionManagementProps> = (props) => {
  const { appId } = useParams();

  const { data } = usePromise(async () => {
    if (!appId) {
      return null;
    }
    return listPermissions({ app_id: Number(appId) });
  });

  if (!data) {
    return loading;
  }

  return data.app_name;
};

export default PermissionManagement;
