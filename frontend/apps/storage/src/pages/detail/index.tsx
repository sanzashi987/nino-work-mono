/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { useCallback, useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import {
  ArrowBack, Check, Close, CloudUpload, CreateNewFolder, HomeRounded
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
  TableHead,
  Badge,
  Menu,
  MenuItem
} from '@mui/material';
import Button from '@mui/material/Button';

import { loading, LoadingGroup, RequestButton, Uploader } from '@nino-work/ui-components';
import { filesize } from 'filesize';
import {
  BucketInfo, createDir, DirInfo, DirResponse, getBucketInfo, listBucketDir,
  uploadFiles
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
  const [draftFolder, setDraftFolder] = useState<boolean>(false);
  const [toUpload, setToUpload] = useState<{ files: File[], map: Record<string, boolean> }>({ files: [], map: {} });
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
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

  const handleCreateDir = useCallback(async () => {
    const name = (ref.current?.firstChild as any)?.value;
    const currentPathId = paths?.at(-1)?.id;
    if (!currentPathId || !name) {
      return Promise.resolve();
    }
    return createDir({ name, bucket_id: Number(id), parent_id: currentPathId })
      .then(() => {
        getBucketDirContent(currentPathId);
        setDraftFolder(false);
      });
  }, [getBucketDirContent, id, paths]);

  const onSelectFile = useCallback((files: File[]) => {
    setToUpload((last) => {
      const toAdd = files.filter((e) => !last.map[e.name]);
      if (toAdd.length === 0) {
        return last;
      }
      const next = { files: last.files.concat(), map: { ...last.map } };
      return toAdd.reduce((l, c) => {
        l.files.push(c);
        // eslint-disable-next-line no-param-reassign
        l.map[c.name] = true;
        return l;
      }, next);
    });
  }, []);

  const openUploadMenu = useCallback((e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    if (toUpload.files.length === 0) {
      return;
    }
    setAnchorEl(e.currentTarget);
  }, [toUpload]);

  const uploadFilesToBucket = useCallback(async () => {
    const currentPathId = paths?.at(-1)?.id;
    if (!currentPathId || toUpload.files.length === 0) {
      return Promise.resolve();
    }
    return uploadFiles({ bucket_id: Number(id), path_id: currentPathId, file: toUpload.files }).then(() => {
      setAnchorEl(null);
      getBucketDirContent(currentPathId);
      setToUpload({ files: [], map: {} });
    });
  }, [id, paths, toUpload]);

  return (
    <Box p={2}>
      <Stack direction="row" alignItems="center" mb={1}>
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
              <IconButton onClick={() => { setDraftFolder(true); }}>
                <CreateNewFolder />
              </IconButton>
              <Uploader onChange={onSelectFile}>
                <IconButton onContextMenu={openUploadMenu}>
                  <Badge badgeContent={toUpload.files.length} color="primary">
                    <CloudUpload />
                  </Badge>
                </IconButton>
              </Uploader>
            </Stack>

            <TableContainer component={Paper} elevation={10}>
              <Table size="small">
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell>Size</TableCell>
                    <TableCell>Update Time</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>

                  {draftFolder
                    && (
                      <TableRow>
                        <TableCell>
                          <Stack direction="row" spacing={1}>
                            <Input ref={ref} size="small" sx={{ input: { height: '12px', fontSize: '12px' } }} />
                            <LoadingGroup>
                              <RequestButton size="small" variant="text" sx={{ minWidth: 0, p: 0 }} onClick={handleCreateDir}>
                                <Check fontSize="small" />
                              </RequestButton>
                              <RequestButton size="small" variant="text" sx={{ minWidth: 0, p: 0 }} onClick={async () => { setDraftFolder(false); }}>
                                <Close fontSize="small" />
                              </RequestButton>
                            </LoadingGroup>
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
                      <TableCell />
                      <TableCell />
                    </TableRow>
                  ))}

                  {dirContents.files.map((e) => (
                    <TableRow key={`f${e.file_id}`}>
                      <TableCell>{e.name}</TableCell>
                      <TableCell>{filesize(e.size)}</TableCell>
                      <TableCell>{e.update_time}</TableCell>
                    </TableRow>
                  ))}

                  {dirContents.dirs.length + dirContents.files.length === 0 && !draftFolder
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

      <Menu
        id="basic-menu"
        anchorEl={anchorEl}
        open={!!anchorEl}
        onClose={() => setAnchorEl(null)}
        MenuListProps={{ 'aria-labelledby': 'basic-button' }}
      >
        {toUpload.files.map((f, i) => (
          <MenuItem
            key={f.name}
            onDoubleClick={() => {
              setToUpload((last) => {
                const next = last.files.concat();
                const nextMap = { ...last.map };
                next.splice(i, 1);
                delete nextMap[f.name];
                return { files: next, map: nextMap };
              });
            }}
          >
            {f.name}
          </MenuItem>
        ))}

        <MenuItem sx={{ minWidth: 200, justifyContent: 'end', px: 1, py: 0 }}>
          <Button fullWidth variant="contained" onClick={uploadFilesToBucket}>
            Upload
          </Button>
        </MenuItem>
      </Menu>

    </Box>
  );
};

export default BucketDetail;
