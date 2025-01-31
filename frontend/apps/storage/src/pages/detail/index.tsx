/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { useCallback, useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import {
  ArrowBack, Check, Close, CreateNewFolder, HomeRounded, UploadFile
} from '@mui/icons-material';
import {
  Stack, IconButton, Typography, Breadcrumbs, Link,
  Table,
  TableRow,
  TableBody,
  TableCell,
  TableContainer,
  Paper,
  Box,
  Input,
  TableHead
} from '@mui/material';
import Button from '@mui/material/Button';
import { loading } from '@nino-work/ui-components';
import {
  BucketInfo, createDir, DirInfo, DirResponse, getBucketInfo, listBucketDir
} from '@/api';

const BucketDetail: React.FC = () => {
  const { id } = useParams();
  if (!id) {
    throw new Error('param id in url is expected');
  }
  const naviagte = useNavigate();

  const [info, setInfo] = useState<BucketInfo | null>(null);
  const [dirContents, setDirContents] = useState<DirResponse | null>(null);
  const [paths, setPaths] = useState<DirInfo[] | undefined>(undefined);
  const [folderDraft, setFolderDraft] = useState<{ pending: boolean } | null>(null);
  const ref = useRef<HTMLInputElement>(null);

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

  const handleCreateDir = useCallback(() => {
    const name = (ref.current?.firstChild as any)?.value;
    const currentPathId = paths?.at(-1)?.id;
    if (!currentPathId || !name) {
      return;
    }
    createDir({ name, bucket_id: Number(id), parent_id: currentPathId })
      .then(() => {
        getBucketDirContent(currentPathId);
        setFolderDraft(null);
      }).catch(() => {
        setFolderDraft({ pending: false });
      });

    setFolderDraft({ pending: true });
  }, [getBucketDirContent, id, paths]);

  return (
    <Box p={2}>
      <Stack direction="row" alignItems="center">
        <IconButton onClick={() => { naviagte('../list'); }}>
          <ArrowBack />
        </IconButton>
        <Typography variant="h5" gutterBottom m={0} ml={1}>
          {info?.code ?? '...'}
        </Typography>
      </Stack>

      {paths
        ? (
          <Stack direction="row" alignItems="center">
            <HomeRounded fontSize="small" />
            <Breadcrumbs maxItems={3}>
              {paths.slice(0, -1).map((p, i) => (
                <Link
                  key={p.id}
                  sx={{ cursor: 'pointer' }}
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
          </Stack>
        ) : null}

      {!dirContents ? loading
        : (
          <>
            <Stack direction="row">
              <IconButton onClick={() => {
                setFolderDraft((last) => ((last === null) ? { pending: false } : last));
              }}
              >
                <CreateNewFolder />
              </IconButton>
              <IconButton>
                <UploadFile />
              </IconButton>
            </Stack>
            <TableContainer component={Paper} elevation={10}>
              <Table size="small">
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>

                  {folderDraft !== null
                    && (
                      <TableRow>
                        <TableCell>
                          <Stack direction="row" spacing={1}>
                            <Input ref={ref} size="small" sx={{ input: { height: '12px', fontSize: '12px' } }} />
                            <Button size="small" variant="text" sx={{ minWidth: 0, p: 0 }} onClick={handleCreateDir} loading={folderDraft.pending}>
                              <Check fontSize="small" />
                            </Button>
                            <Button size="small" variant="text" sx={{ minWidth: 0, p: 0 }} onClick={() => { setFolderDraft(null); }} loading={folderDraft.pending}>
                              <Close fontSize="small" />
                            </Button>
                          </Stack>
                        </TableCell>
                      </TableRow>
                    )}

                  {dirContents.dirs.map((e) => (
                    <TableRow key={`d${e.id}`}>
                      <TableCell>
                        <Link
                          underline="hover"
                          sx={{ cursor: 'pointer' }}
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

                  {dirContents.dirs.length + dirContents.files.length === 0 && folderDraft === null
                && (
                  <Stack justifyContent="center" textAlign="center" minHeight="200px">
                    No Data
                  </Stack>
                )}
                </TableBody>
              </Table>
            </TableContainer>
          </>

        )}

    </Box>
  );
};

export default BucketDetail;
