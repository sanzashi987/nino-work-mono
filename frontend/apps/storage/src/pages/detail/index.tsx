/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { useCallback, useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { ArrowBack } from '@mui/icons-material';
import {
  Stack, IconButton, Typography, Breadcrumbs, Link,
  Table,
  TableRow,
  TableBody,
  TableCell,
  TableContainer,
  Paper,
  Box
} from '@mui/material';
import { loading } from '@nino-work/ui-components';
import {
  BucketInfo, DirInfo, DirResponse, getBucketInfo, listBucketDir
} from '@/api';

type BucketDetailProps = {};

const BucketDetail: React.FC<BucketDetailProps> = (props) => {
  const { id } = useParams();
  if (!id) {
    throw new Error('param id in url is expected');
  }
  const naviagte = useNavigate();

  const [info, setInfo] = useState<BucketInfo | null>(null);
  const [dirContents, setDirContents] = useState<DirResponse | null>(null);
  const [paths, setPaths] = useState<DirInfo[] | undefined>(undefined);

  useEffect(() => {
    getBucketInfo({ bucket_id: id }).then((res) => {
      setInfo(res);
      setDirContents(res.dir_contents);
      setPaths([{ id: res.root_path_id, name: res.code }]);
    });
  }, []);

  const getBucketDirContent = useCallback((pathId: number) => {
    setDirContents(null);
    listBucketDir({ bucket_id: id, path_id: pathId }).then(setDirContents);
  }, []);

  return (
    <Box p={2}>
      <Stack direction="row" alignItems="center">
        <IconButton onClick={() => { naviagte('..'); }}>
          <ArrowBack />
        </IconButton>
        <Typography variant="h5" gutterBottom m={0} ml={1}>
          {info?.code ?? '...'}
        </Typography>
      </Stack>

      {paths
        ? (
          <Breadcrumbs maxItems={3}>
            {paths.slice(0, -1).map((p, i) => (
              <Link
                key={p.id}
                underline="hover"
                color="inherit"
                onClick={() => {
                  setPaths((last) => last?.slice(0, i + 1));
                  getBucketDirContent(p.id);
                }}
              >
                {p.name}
              </Link>
            ))}
            <Typography sx={{ color: 'text.primary' }}>
              {paths.at(-1)?.name}
            </Typography>
          </Breadcrumbs>
        ) : null}

      {!dirContents ? loading
        : (
          <TableContainer component={Paper} elevation={10} sx={{ mt: 2 }}>
            <Table>
              <TableBody>
                {dirContents.dirs.map((e) => (
                  <TableRow key={`d${e.id}`}>
                    <TableCell>
                      <Link
                        underline="hover"
                        onClick={() => {
                          setPaths((last) => last?.concat(e));
                          getBucketDirContent(e.id);
                        }}
                      >
                        {e.name}
                      </Link>
                    </TableCell>
                  </TableRow>
                ))}
                {dirContents.files.map((e) => (
                  <TableRow key={`f${e.file_id}`}>
                    <TableCell>{e.name}</TableCell>
                    <TableCell>{e.update_time}</TableCell>
                  </TableRow>
                ))}

                {dirContents.dirs.length + dirContents.files.length === 0
                && (
                  <Stack justifyContent="center" textAlign="center" minHeight="200px">
                    No Data
                  </Stack>
                )}
              </TableBody>
            </Table>
          </TableContainer>
        )}

    </Box>
  );
};

export default BucketDetail;
