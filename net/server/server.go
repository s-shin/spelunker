package server

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/boltdb/bolt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/s-shin/spelunker/net/shogi"
)

type Server struct {
	db *bolt.DB
}

func NewServer() (*Server, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	if err := initDB(db); err != nil {
		return nil, err
	}
	return &Server{
		db: db,
	}, nil
}

func initDB(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Server) Close() {
	s.db.Close()
}

func (s *Server) Handshake(ctx context.Context, req *pb.HandshakeRequest) (*pb.HandshakeResult, error) {
	switch req.GetVersion() {
	case 1:
	default:
		return nil, grpc.Errorf(codes.Unimplemented, "not supported version")
	}
	return &pb.HandshakeResult{}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponce, error) {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		if b == nil {
			return grpc.Errorf(codes.Internal, "")
		}
		data := b.Get([]byte(req.GetUsername()))
		if data != nil {
			if err := bcrypt.CompareHashAndPassword(data, []byte(req.GetPassword())); err != nil {
				return grpc.Errorf(codes.InvalidArgument, "invalid password")
			}
			return nil
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		b.Put([]byte(req.GetUsername()), hash)
		b.Put([]byte(req.GetUsername()), []byte(req.GetPlayer().String()))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponce, error) {
	return nil, nil
}
