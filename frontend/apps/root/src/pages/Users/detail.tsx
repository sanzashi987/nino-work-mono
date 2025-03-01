import { Enum } from '@nino-work/shared';
import { OpenModalContext } from '@nino-work/ui-components';
import React, { useContext, useState } from 'react';

type UserDetailProps = {};

const UserDetail: React.FC<UserDetailProps> = (props) => {
  const { } = useContext(OpenModalContext);
  const [roles, setRoles] = useState<Enum<number>[]>([]);
};

const openUserDetail = () => { };
