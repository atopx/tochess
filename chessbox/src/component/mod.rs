use bevy::{prelude::{Resource, Component, States}, time::{Timer, TimerMode}};


/// 棋子
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Pieces {
    /// 车
    R,
    /// 马
    N,
    /// 象
    B,
    /// 士
    A,
    /// 将
    K,
    /// 跑
    C,
    /// 兵
    P,
}

/// 棋子颜色
#[derive(Default, Debug, Clone, Copy, PartialEq, Eq, Hash, States)]
pub enum PieceColor {
    #[default]
    Red,
    Black,
}


/// 棋子位置
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Position {
    pub row: i8,
    pub col: i8,
}

impl Position {
    pub const fn new(row: i8, col: i8) -> Self {
        Self { row, col }
    }
}

/// Piece 棋子
#[derive(Debug, Clone, Copy, Component, PartialEq, Eq)]
pub struct Piece {
    pub pos: Position,
    pub amount_moved: u32,
    pub piece_type: Pieces,
    pub color: PieceColor,
}

/// 分数显示组件
#[derive(Component)]
pub struct DisplayScore;

#[derive(Component)]
pub struct OptionAI;

#[derive(Component)]
pub struct OptionAI;
